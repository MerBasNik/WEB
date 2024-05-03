import React, { useEffect, useState } from "react";
import styles from "./homePage.module.css";
import { ManInChairIcon } from "../../../components/icons/icons";
import { MeetingForm } from "../../../components/formForMeet/meetingForm";
import { MeetCard } from "../../../components/meetCard/meetCard";

export const HomePage = () => {
	const [activeForm, setActiveForm] = useState(false);
	const [chats, setChats] = useState<{ id: number; title: string }[]>([]);

	let classesForIcon = styles.mainContent__iconMan;
	let classesForForm =
		styles.mainContent__form + " " + styles.wrapper__hidden;
	if (activeForm) {
		classesForIcon += " " + styles.wrapper__hidden;
		classesForForm = styles.mainContent__form;
	} else {
		classesForIcon = styles.mainContent__iconMan;
		classesForForm += " " + styles.wrapper__hidden;
	}

	const getChats = async () => {
		try {
			const res = await fetch(
				"http://localhost:8000/api/chats//get_all_chats",
				{
					method: "GET",
				},
			);
			const data = await res.json();
			if (res.ok) {
				setChats(data);
			}
		} catch (err) {
			console.log(err);
		}
	};

	useEffect(() => {
		getChats();
	}, []);

	return (
		<div className={styles.homePage}>
			<div className={styles.homePage__header}>
				<div className={styles.header__text}>
					<h1 className={styles.text}>Ваши встречи</h1>
				</div>
				<div className={styles.header__buttonOfMeetings}>
					<button
						className={styles.buttonOfMeetings}
						onClick={() => {
							setActiveForm(!activeForm);
						}}
					>
						Назначить встречу
					</button>
				</div>
			</div>
			<div className={styles.homePage__mainContent}>
				<div className={classesForIcon}>
					<ManInChairIcon />
				</div>
				<div className={styles.mainContent__meetings}>
					{/* {chats.map((chat, index) => (
						<MeetCard
							chatId={chat.id}
							title={chat.title}
							key={index}
						/>
					))} */}
				</div>
				<div className={classesForForm}>
					<MeetingForm
						active={activeForm}
						setActive={setActiveForm}
					/>
				</div>
			</div>
		</div>
	);
};
