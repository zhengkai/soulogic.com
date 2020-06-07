import { pb } from '../pb/pb';
import { Api } from './api';

export class ApiPublic {

	static async blogIndex() {
		const re = await this._single('blogIndex', true);
		return re;
	}

	static async _single(k: string, v: unknown) {
		const req = { } as pb.ReqPublic;
		req[k] = v;

		const list = await this._gateway([req]);

		const re = (list?.[0]?.[k]) || null;
		return re;
	}

	static async _gateway(list: pb.ReqPublic[]) {

		const o = {
			ts: Date.now(),
		} as pb.ReqGateway;

		o.public = {
			list,
		} as pb.IReqPublicList;

		const re = await Api.gateway(o);

		const ts = re?.ts;
		if (!ts) {
			console.warn('server time fail');
			return;
		}

		const op = (re?.public?.list) || null;

		return op;
	}
}
