export namespace codebox {
	
	export class KeyPair {
	    PrivateKey: string;
	    PublicKey: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyPair(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.PrivateKey = source["PrivateKey"];
	        this.PublicKey = source["PublicKey"];
	    }
	}

}

