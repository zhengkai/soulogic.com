import * as marked from 'marked';
import hljs from 'highlight.js';

export class Api {

	static init = false;

	static re: string[] = [];

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
