import React from "react";
import styles from "./meetCard.module.css";
import { LineSeparatorMeetCard } from "../icons/icons";

type MeetCardProps = {
	chatId: number;
	username: string;
	userId: number;
	title: string;
};

export const MeetCard = (props: MeetCardProps) => {
	const { chatId, title, userId, username } = props;

	const joinChat = (chatId: string) => {
		const ws = new WebSocket(
			`http://localhost:8000/api/ws/joinRoom/${chatId}?userId=${userId}&username=${username}`,
		);
		// if (ws.OPEN) {
		// 	setConn(ws);
		// 	router.push("/app");
		// 	return;
		// }
	};

	return (
		<div className={styles.wrapper__meetCard}>
			<div className={styles.wrapper__users}></div>
			<div className={styles.wrapper__separator}>
				<LineSeparatorMeetCard />
			</div>
			<div className={styles.wrapper__text}>
				<h1 className={styles.text}>{title}</h1>
			</div>
			<div className={styles.wrapper__separator}>
				<LineSeparatorMeetCard />
			</div>
			<div className={styles.wrapper__activities}>
				<h2>{chatId}</h2>
			</div>
		</div>
	);
};
