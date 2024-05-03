import React from "react";
import stylesWelcome from "./welcome.module.css";
import stylesContacts from "../contacts/contactsPage.module.css";
import { Button, ButtonType } from "../../../components/button/button";
import { useNavigate } from "react-router-dom";
import {
	StageType,
	UserPath,
} from "../../../components/userPathLeftBlock/userPath";

export const WelcomePage = () => {
	const navigate = useNavigate();

	return (
		<div className={stylesContacts.wrapper__contacts}>
			<div className={stylesContacts.wrapper__userPath}>
				{<UserPath Stage={StageType.Third} />}
			</div>
			<div className={stylesWelcome.wrapper__welcomeBlock}>
				<form
					onSubmit={(event) => {
						navigate("/main/home", {
							replace: true,
						});
					}}
				>
					<div className={stylesWelcome.wrapper__appeal}>
						<div className={stylesWelcome.wrapper__welcomeHeader}>
							<h1 className={stylesWelcome.welcomeHeader}>
								Добро пожаловать!
							</h1>
						</div>
						<div className={stylesWelcome.wrapper__welcomeText}>
							<p className={stylesWelcome.welcomeText}>
								Чтобы найти собеседника на чашку кофе,
							</p>
							<p className={stylesWelcome.welcomeText}>
								нам важно, чтобы ваши интересы
							</p>
							<p className={stylesWelcome.welcomeText}>
								совпадали. Для качественного поиска нам
							</p>
							<p className={stylesWelcome.welcomeText}>
								нужна информация о вас.
							</p>
						</div>
					</div>
					<div className={stylesContacts.profile__footer}>
						<div className={stylesContacts.wrapper__buttonNext}>
							<Button
								id={"welcome"}
								title={"Далее"}
								type={ButtonType.Text}
								typeButton={"submit"}
							/>
						</div>
					</div>
				</form>
			</div>
		</div>
	);
};
