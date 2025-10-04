export namespace request {
	
	export class CreateConnection {
	    host: string;
	    port: string;
	    database_name: string;
	    user: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateConnection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database_name = source["database_name"];
	        this.user = source["user"];
	        this.password = source["password"];
	    }
	}

}

