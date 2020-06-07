import { pb } from '../pb/pb';
import { Api } from './api';

export class ApiOp {

	static async echo(s: string) {
		const re = await this._single('echo', s);
		return re;
	}

	static async itemEdit(id: number, raw: string, format: pb.PostFormat.Enum = pb.PostFormat.Enum.Markdown) {

		const ie = {
			ID: id,
			revision: {
				format,
				raw,
			} as pb.Revision,
		} as pb.IItemEdit;

		const re = await this._single('itemEdit', ie);
		return re;
	}

	static async _single(k: string, v: unknown) {
		const req = { } as pb.ReqOp;
		req[k] = v;

		const list = await this._gateway([req]);

		const re = (list?.[0]?.[k]) || null;
		return re;
	}

	static async _gateway(list: pb.ReqOp[]) {

		const o = {
			ts: Date.now(),
		} as pb.ReqGateway;

		o.op = {
			auth: {
				token: 'token',
			} as pb.Auth,
			list,
		} as pb.IReqOpList;

		const re = await Api.gateway(o);

		const ts = re?.ts;
		if (!ts) {
			console.warn('server time fail');
			return;
		}

		const op = (re?.op?.list) || null;

		return op;
	}
}
