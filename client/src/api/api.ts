import * as marked from 'marked';
import hljs from 'highlight.js';

export class Api {

	static init = false;

	static async fetch() {
		const re: string[] = [];
		for (let i = 1; i <= 9; i++) {
			const text = await this.fetchOne(i);
			re.push(text);
		}
		return re;
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
