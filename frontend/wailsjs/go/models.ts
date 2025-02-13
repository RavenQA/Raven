export namespace types {
	
	export class BrowserListItem {
	    Version: string;
	    Name: string;
	    ReleaseDate: string;
	    Path: string;
	    Available: boolean;
	
	    static createFrom(source: any = {}) {
	        return new BrowserListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Name = source["Name"];
	        this.ReleaseDate = source["ReleaseDate"];
	        this.Path = source["Path"];
	        this.Available = source["Available"];
	    }
	}

}

