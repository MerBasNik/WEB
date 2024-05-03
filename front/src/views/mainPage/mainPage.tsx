import React from "react";
import { Outlet } from "react-router-dom";
import styles from "./mainPage.module.css";
import {
	ButtonTextLink,
	LocationOnPages,
} from "../../components/button/button";

export const MainPage = () => {
	return (
		<div className={styles.mainPage}>
			<header className={styles.mainPage__header}>
				<div className={styles.header__buttons}>
					<div className={styles.header__buttonHome}>
						<ButtonTextLink
							location={LocationOnPages.HomePage}
							id={"home"}
							link={"/main/home"}
							title={"Главная"}
						/>
					</div>
					<div className={styles.header__buttonProfile}>
						<ButtonTextLink
							location={LocationOnPages.HomePage}
							id={"profile"}
							link={"/main/profile"}
							title={"Профиль"}
						/>
					</div>
				</div>
			</header>
			<div className={styles.mainContent}>
				<Outlet />
			</div>
		</div>
	);
};
