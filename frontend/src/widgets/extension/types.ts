export interface AccountI {
	domain: string;
	username: string;
	password: string;
	note: string;
}

export interface LinkI {
	linkPage?: string;
	targetUrl?: string;
	anchor?: string;
	rel?: string;
	reviewedAt?: Date;
	stetus?: 'active' | 'inactive' | 'pending';
}
