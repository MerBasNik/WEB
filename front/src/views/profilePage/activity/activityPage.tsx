import React, { useState } from "react";
import {
	Button,
	ButtonTextLink,
	ButtonType,
	LocationOnPages,
} from "../../../components/button/button";
import stylesContacts from "../contacts/contactsPage.module.css";
import stylesActivity from "./activityPage.module.css";
import { ChoiceActivity } from "../../../components/choiceActivity/choiceActivity";
import { AllActivities } from "../../../data/data";
import { useLocalStorage } from "../../../hooks/useLocalStorage";
import { useNavigate } from "react-router-dom";
import { TypeOfActivities } from "../../../data/types";
import {
	StageType,
	UserPath,
} from "../../../components/userPathLeftBlock/userPath";

export const ActivityPage = () => {
	const [description, setActivities] = useState(Array<TypeOfActivities>);

	const [token, setToken] = useLocalStorage({
		initialValue: "",
		key: "token",
	});

	const [profile_id, setProfileId] = useLocalStorage({
		initialValue: 0,
		key: "profile_id",
	});

	const navigate = useNavigate();

	return (
		<div className={stylesContacts.wrapper__contacts}>
			<div className={stylesContacts.wrapper__userPath}>
				{<UserPath Stage={StageType.Second} />}
			</div>
			<div className={stylesActivity.wrapper__activities}>
				<div className={stylesContacts.contacts__mainHeader}>
					<h1 className={stylesContacts.mainHeader}>
						Выберите то, что вас описывает
					</h1>
				</div>
				<div className={stylesContacts.contacts__header__h2}>
					<h2 className={stylesContacts.header__h2}>
						Это нужно чтобы вам было, что обсудить на встрече
					</h2>
				</div>
				{/* <div> Поиск </div> */}
				<form
					onSubmit={async (event) => {
						event.preventDefault();

						const data = {
							description,
						};

						const response = await fetch(
							`http://localhost:8000/api/profile/${profile_id}/hobby/create_hobby`,
							{
								method: "POST",
								body: JSON.stringify(data),
								headers: {
									"Content-Type": "application/json",
									Authorization: `Bearer ${token}`,
									prof_id: profile_id,
								},
								credentials: "include",
							},
						);
						if (!response.ok) {
							console.log(response.json());
							return;
						}
						navigate("/auth_profile/welcome", {
							replace: true,
						});
					}}
					action="http://localhost:8000/api/profile/create_profile"
					method="post"
				>
					<div className={stylesActivity.wrapper__choiceActivity}>
						{
							<ChoiceActivity
								AllActivities={AllActivities}
								activities={description}
								setActivities={setActivities}
							/>
						}
					</div>

					<div className={stylesContacts.profile__footer}>
						<div className={stylesContacts.wrapper__buttonNext}>
							<ButtonTextLink
								location={LocationOnPages.HomePage}
								id={"back"}
								title={"Назад"}
								link={"/auth_profile/contacts"}
							/>
						</div>
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
