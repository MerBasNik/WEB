import { GroupActivity, Headers, TypeAllActivities } from "./types";

export const Trends: GroupActivity = {
	header: Headers.Trends,
	activities: [
		{ description: "NFL" },
		{ description: "NBA" },
		{ description: "Мировые новости" },
		{ description: "ChatGPT" },
		{ description: "One Piece" },
		{ description: "Midjourney" },
		{ description: "Cплетни о знаменитостях" },
	],
};

export const Games: GroupActivity = {
	header: Headers.Games,
	activities: [
		{ description: "Call of Duty" },
		{ description: "Baldur’s Gate 3" },
		{ description: "Minecraft" },
		{ description: "Playstation" },
		{ description: "Genshin Impact" },
		{ description: "GTA" },
		{ description: "Sims" },
		{ description: "Terraria" },
		{ description: "Red Dead Redemption" },
	],
};

export const AllActivities: TypeAllActivities = {
	allActivities: [Trends, Games],
};
