import React, { useEffect, useState } from "react";
import styles from "./userPath.module.css";
import { Link, useLocation } from "react-router-dom";
import { LineSeparator } from "../icons/icons";

export enum StageType {
	First,
	Second,
	Third,
}

type UserPathProps = {
	Stage: StageType;
};

export const UserPath = (props: UserPathProps) => {
	const { Stage } = props;

	let classesForFirstElList = styles.list__elementOfList;
	let classesForSecondElList = styles.list__elementOfList;
	let classesForThirdElList = styles.list__elementOfList;

	if (Stage === StageType.First) {
		classesForFirstElList += " " + styles.elementOfList__active;
	}

	if (Stage === StageType.Second) {
		classesForSecondElList += " " + styles.elementOfList__active;
	}

	if (Stage === StageType.Third) {
		classesForThirdElList += " " + styles.elementOfList__active;
	}

	return (
		<div className={styles.wrapper__userPath}>
			<div className={styles.userPath__list}>
				<ol className={styles.list}>
					<li className={classesForFirstElList}>
						Персональная информация
					</li>
					<div className={styles.lineSeparator}>
						{<LineSeparator />}
					</div>
					<li className={classesForSecondElList}>Ваши интересы</li>
					<div className={styles.lineSeparator}>
						{<LineSeparator />}
					</div>
					<li className={classesForThirdElList}>Добро пожаловать!</li>
				</ol>
			</div>
		</div>
	);
};
