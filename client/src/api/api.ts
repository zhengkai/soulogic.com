import * as marked from 'marked';
import { hljs } from 'highlight.js';
import { pb } from '../pb/pb';

export class Api {

	static init = false;

	static url = '/gateway';

	static re: string[] = [];

	static async post(body: Uint8Array) {

		const res = await fetch(this.url, {
			method: 'POST',
			body,
		});

		if (!res.ok) {
			console.warn(`server error, HTTP Status ${res.status}`);
			return;
		}

		const buf = await res.arrayBuffer();

		return new Uint8Array(buf);
	}

	static async gateway(req: pb.ReqGateway) {

		const ab = pb.ReqGateway.encode(req).finish();

		const re = await this.post(ab);
		if (!re) {
			console.warn(`post fail, req: ${req}`);
			return;
		}
		if (!(re as Uint8Array).length) {
			console.warn('server error, response empty');
			return;
		}

		const x = pb.RspGateway.decode(re as Uint8Array);
		if (!x) {
			console.warn('decode RspGateway fail');
			return;
		}

		return x;
	}

	static async fetch() {

		if (this.re.length > 0) {
			return this.re;
		}

		for (let i = 1; i <= 9; i++) {
			const text = await this.fetchOne(i);
			this.re.push(text);
		}
		return this.re;
	}

	static async fetchOne(i: number) {

		let s = '' + i;
		if (i < 10) {
			s = '0' + s;
		}

		const r = await fetch(`/assets/markdown/${s}.txt`);
		const text = await r.text();

		return text;
	}

	static markdown(markdown: string) {
		if (!this.init) {
			this.init = true;
			hljs.configure({ tabReplace: '__', });
		}
		return marked(markdown, {
			gfm: true,
			highlight: (code, lang) => {
				if (!lang) {
					return code;
				}
				return hljs.highlight(lang, code).value;
			},
		});
	}
}
