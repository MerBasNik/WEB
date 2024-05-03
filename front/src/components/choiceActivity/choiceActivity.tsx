import React, { useState } from "react";
import styles from "./choiceActivity.module.css";
import { Headers, TypeAllActivities, TypeOfActivities } from "../../data/types";
import {
	BackButtonIcon,
	ComputerGamesActivity,
	NextButtonIcon,
	TrendsActivity,
} from "../icons/icons";
import { Activity } from "../activity/activity";
import { Button, ButtonType } from "../button/button";

type ChoiceActivityProps = {
	AllActivities: TypeAllActivities;
	activities: Array<TypeOfActivities>;
	setActivities: (activities: Array<TypeOfActivities>) => void;
};

function ChoiceHeaderAndIcon(currentTitle: Headers) {
	let currentIcon: JSX.Element;
	let headerText: string;
	switch (currentTitle) {
		case Headers.Trends:
			currentIcon = <TrendsActivity />;
			headerText = "Тренды";
			break;
		case Headers.Games:
			currentIcon = <ComputerGamesActivity />;
			headerText = "Компьютерные игры";
			break;
	}
	return { currentIcon, headerText };
}

export const ChoiceActivity = (props: ChoiceActivityProps) => {
	const { AllActivities, activities, setActivities } = props;
	const [index, setIndex] = useState(0);
	const currentTitle = AllActivities.allActivities[index].header;
	const currentArrayActivities =
		AllActivities.allActivities[index].activities;

	const { currentIcon, headerText } = ChoiceHeaderAndIcon(currentTitle);

	return (
		<div className={styles.wrapper__activity}>
			<div className={styles.wrapper__Blockheader}>
				<div className={styles.wrapper__iconActivity}>
					{currentIcon}
				</div>
				<div className={styles.wrapper__Textheader}>
					<h3 className={styles.header__text}>{headerText}</h3>
				</div>
			</div>
			<div className={styles.wrapper__activities}>
				{currentArrayActivities.map((elem) => (
					<Activity
						title={elem}
						key={elem.description}
						activities={activities}
						setActivities={setActivities}
					/>
				))}
			</div>
			<div className={styles.wrapper__arrows}>
				<div className={styles.wrapper__arrowBack}>
					<Button
						id={"back"}
						type={ButtonType.Icon}
						icon={<BackButtonIcon />}
						onClick={() => {
							if (index - 1 >= 0) {
								setIndex(index - 1);
							}
						}}
					/>
				</div>
				<div className={styles.wrapper__arrowNext}>
					<Button
						id={"back"}
						type={ButtonType.Icon}
						icon={<NextButtonIcon />}
						onClick={() => {
							if (
								index + 1 <
								AllActivities.allActivities.length
							) {
								setIndex(index + 1);
							}
						}}
					/>
				</div>
			</div>
		</div>
	);
};
