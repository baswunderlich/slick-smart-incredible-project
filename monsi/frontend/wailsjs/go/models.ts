export namespace util {
	
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
	export class Proof {
	    type: string;
	    proofValue: string;
	
	    static createFrom(source: any = {}) {
	        return new Proof(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.proofValue = source["proofValue"];
	    }
	}
	export class VC {
	    "@context": string[];
	    id: string;
	    type: string[];
	    issuer: string;
	    validFrom: string;
	    validUntil: string;
	    credentialSubject: {[key: string]: any};
	    proof: Proof;
	
	    static createFrom(source: any = {}) {
	        return new VC(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["@context"] = source["@context"];
	        this.id = source["id"];
	        this.type = source["type"];
	        this.issuer = source["issuer"];
	        this.validFrom = source["validFrom"];
	        this.validUntil = source["validUntil"];
	        this.credentialSubject = source["credentialSubject"];
	        this.proof = this.convertValues(source["proof"], Proof);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

