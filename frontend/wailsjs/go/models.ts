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
	export class Credential {
	    id: number;
	    title: string;
	    db_vendor: string;
	    hex_color: string;
	    host: string;
	    port: string;
	    user: string;
	    password: string;
	    database_name: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Credential(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.db_vendor = source["db_vendor"];
	        this.hex_color = source["hex_color"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.user = source["user"];
	        this.password = source["password"];
	        this.database_name = source["database_name"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
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
	export class DeleteCredential {
	    credential_id: number;
	
	    static createFrom(source: any = {}) {
	        return new DeleteCredential(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.credential_id = source["credential_id"];
	    }
	}
	export class GetCredentials {
	    search: string;
	
	    static createFrom(source: any = {}) {
	        return new GetCredentials(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.search = source["search"];
	    }
	}
	export class GetTableColumns {
	    connection_id: string;
	    table_name: string;
	
	    static createFrom(source: any = {}) {
	        return new GetTableColumns(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	        this.table_name = source["table_name"];
	    }
	}
	export class GetTableRecords {
	    connection_id: string;
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
	        this.connection_id = source["connection_id"];
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
	export class GetTables {
	    connection_id: string;
	
	    static createFrom(source: any = {}) {
	        return new GetTables(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection_id = source["connection_id"];
	    }
	}
	export class SaveCredential {
	    id: number;
	    title: string;
	    db_vendor: string;
	    host: string;
	    port: string;
	    user: string;
	    password: string;
	    database_name: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new SaveCredential(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.db_vendor = source["db_vendor"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.user = source["user"];
	        this.password = source["password"];
	        this.database_name = source["database_name"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}

}

export namespace response {
	
	export class CreateConnection {
	    ConnectionId: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateConnection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ConnectionId = source["ConnectionId"];
	    }
	}
	export class GetCredentials {
	    Data: model.Credential[];
	    Total: number;
	
	    static createFrom(source: any = {}) {
	        return new GetCredentials(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = this.convertValues(source["Data"], model.Credential);
	        this.Total = source["Total"];
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

