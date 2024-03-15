import React from "react";
import styles from "./mainPage.module.css";
import { ButtonIconLink } from "../../components/button/button";
import {
	ChatsMainPageIcon,
	CoffeeMainPageIcon,
	ProfileImageIcon,
	ProfileMainPageIcon,
} from "../../components/icons/icons";
import { Outlet, Route, Routes } from "react-router-dom";
import { Profile } from "../profilePage/profile";

export const MainPage = () => {
	return (
		<div>
			<div className={styles.leftPanel}>
				<div className={styles.leftPanel__profileIcon}>
					<ProfileImageIcon />
				</div>
				<div className={styles.leftPanel__pagesIcon}>
					<ButtonIconLink
						id="mainPage"
						link={"/main"}
						icon={<CoffeeMainPageIcon />}
					/>
					<ButtonIconLink
						id="chats"
						link={"/main/chats"}
						icon={<ChatsMainPageIcon />}
					/>
					<ButtonIconLink
						id="profile"
						link={"/main/profile"}
						icon={<ProfileMainPageIcon />}
					/>
				</div>
			</div>
			<div className={styles.rightPanel}>
				<Outlet />
			</div>
		</div>
	);
};
