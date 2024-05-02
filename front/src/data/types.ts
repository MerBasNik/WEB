export enum Headers {
	Trends,
	Games,
}

export type TypeOfActivities = {
	description: string;
};

export type GroupActivity = {
	header: Headers;
	activities: Array<TypeOfActivities>;
};

export type TypeAllActivities = {
	allActivities: Array<GroupActivity>;
};
