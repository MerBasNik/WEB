import React, { useMemo, useState } from "react";
import styles from "./profile.module.css";

type InputBlock = {
	title: string;
};

const InputBlock = (props: InputBlock) => {
	const { title } = props;
	const [inputText, setInputText] = useState("");

	let isSpanTextBig = "";

	if (inputText.length < 30) {
		isSpanTextBig = inputText;
	}

	const saveSpanText = useMemo(() => {
		return inputText;
	}, [isSpanTextBig]);

	return (
		<div className={styles.resize__container}>
			<span className={styles.resize__text}>{saveSpanText}</span>
			<input
				className={styles.resize__input}
				value={inputText}
				placeholder={title}
				onChange={(event) => {
					setInputText(event.target.value);
				}}
			/>
		</div>
	);
};

export const Profile = () => {
	return (
		<div className={styles.mainPage}>
			<div className={styles.profile}>
				<div
					className={`${styles.profile__Level} ${styles.border__bottom__level}`}
				>
					<div className={styles.wrapper__profile__icon}>
						<button className={styles.profile__icon}></button>
					</div>
					<div className={styles.wrapper__profile__initials}>
						<InputBlock title={"Мое имя"} />
						<InputBlock title={"Моя фамилия"} />
					</div>
				</div>
				<div
					className={`${styles.profile__Level} ${styles.border__bottom__level}`}
				>
					<InputBlock title={"Город"} />
					<InputBlock title={"Профиль в телеграмме"} />
				</div>
				<div
					className={`${styles.profile__Level} ${styles.border__bottom__level}`}
				>
					<InputBlock title={"Хобби"} />
				</div>
				<div className={styles.profile__Level}></div>
			</div>
		</div>
	);
};
