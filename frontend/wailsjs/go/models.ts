export namespace model {
	
	export class Condition {
	    field: string;
	    operator: string;
	    first_value: string;
	    second_value: string;
	
	    static createFrom(source: any = {}) {
	        return new Condition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.field = source["field"];
	        this.operator = source["operator"];
	        this.first_value = source["first_value"];
	        this.second_value = source["second_value"];
	    }
	}

}

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
	export class GetTableRecords {
	    table_name: string;
	    limit: number;
	    offset: number;
	    sort_by: string;
	    order_by: string;
	    conditions: model.Condition[];
	
	    static createFrom(source: any = {}) {
	        return new GetTableRecords(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.table_name = source["table_name"];
	        this.limit = source["limit"];
	        this.offset = source["offset"];
	        this.sort_by = source["sort_by"];
	        this.order_by = source["order_by"];
	        this.conditions = this.convertValues(source["conditions"], model.Condition);
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

export namespace response {
	
	export class GetTableRecords {
	    Data: any[];
	    Total: number;
	    Limit: number;
	    Offset: number;
	
	    static createFrom(source: any = {}) {
	        return new GetTableRecords(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = source["Data"];
	        this.Total = source["Total"];
	        this.Limit = source["Limit"];
	        this.Offset = source["Offset"];
	    }
	}

}

