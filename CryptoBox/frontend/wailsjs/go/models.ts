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
	export class SHA1MD5Result {
	    sha1: string;
	    md5: string;
	
	    static createFrom(source: any = {}) {
	        return new SHA1MD5Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.sha1 = source["sha1"];
	        this.md5 = source["md5"];
	    }
	}
	export class SHA2Result {
	    sha224: string;
	    sha256: string;
	    sha384: string;
	    sha512: string;
	
	    static createFrom(source: any = {}) {
	        return new SHA2Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.sha224 = source["sha224"];
	        this.sha256 = source["sha256"];
	        this.sha384 = source["sha384"];
	        this.sha512 = source["sha512"];
	    }
	}
	export class SHA3Result {
	    sha3_224: string;
	    sha3_256: string;
	    sha3_384: string;
	    sha3_512: string;
	
	    static createFrom(source: any = {}) {
	        return new SHA3Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.sha3_224 = source["sha3_224"];
	        this.sha3_256 = source["sha3_256"];
	        this.sha3_384 = source["sha3_384"];
	        this.sha3_512 = source["sha3_512"];
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

