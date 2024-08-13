export namespace wallet {
	
	export class DID {
	    did: string;
	    pubKey: string;
	    privKey: string;
	
	    static createFrom(source: any = {}) {
	        return new DID(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.did = source["did"];
	        this.pubKey = source["pubKey"];
	        this.privKey = source["privKey"];
	    }
	}

}

