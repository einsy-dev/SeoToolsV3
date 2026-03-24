export namespace extension {
	
	export class State {
	    Links: string[];
	    Index: number;
	    Current: string;
	    CurrentDomain: string;
	
	    static createFrom(source: any = {}) {
	        return new State(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Links = source["Links"];
	        this.Index = source["Index"];
	        this.Current = source["Current"];
	        this.CurrentDomain = source["CurrentDomain"];
	    }
	}

}

