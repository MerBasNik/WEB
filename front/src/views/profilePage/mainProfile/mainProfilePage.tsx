import React from "react";
import styles from "./mainProfilePage.module.css";
import { Outlet } from "react-router-dom";
import { UserPath } from "../../../components/userPathLeftBlock/userPath";
import {
	BackgroundOnProfileIcon,
	ProfileManLeftIcon,
	ProfileManRightIcon,
} from "../../../components/icons/icons";

export const MainProfile = () => {
	return (
		<div className={styles.wrapper__fillProfile}>
			<div className={styles.backrgoundMan__left}>
				<ProfileManLeftIcon />
			</div>
			<div className={styles.background}>
				<div className={styles.background__icon}>
					<BackgroundOnProfileIcon />
				</div>
				<div className={styles.wrapper__auth}>
					<Outlet />
				</div>
			</div>
			<div className={styles.backrgoundMan__right}>
				<ProfileManRightIcon />
			</div>
		</div>
	);
};
