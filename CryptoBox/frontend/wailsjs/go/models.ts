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
	export class SignResult {
	    signature: string;
	    errmsg: string;
	
	    static createFrom(source: any = {}) {
	        return new SignResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.signature = source["signature"];
	        this.errmsg = source["errmsg"];
	    }
	}
	export class SignVerifyResult {
	    isvalid: boolean;
	    errmsg: string;
	
	    static createFrom(source: any = {}) {
	        return new SignVerifyResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.isvalid = source["isvalid"];
	        this.errmsg = source["errmsg"];
	    }
	}

}

