import React, { useState } from "react";
import styles from "./contactsPage.module.css";
import { Button, ButtonType } from "../../../components/button/button";
import { ProfileIcon } from "../../../components/icons/icons";
import {
	LocationInputField,
	TextField,
	TypeInputOnProfile,
} from "../../../components/input/input";
import { useNavigate } from "react-router-dom";
import { useLocalStorage } from "../../../hooks/useLocalStorage";
import {
	StageType,
	UserPath,
} from "../../../components/userPathLeftBlock/userPath";

export const UserContacts = () => {
	const [nameUser, setNameUser] = useState("");
	const [nameUserError, setNameUserError] = useState(false);
	const [surnameUser, setSurNameUser] = useState("");
	const [surnameUserError, setSurnameUserError] = useState(false);
	const [emailUser, setEmailUser] = useState("");
	const [emailUserError, setEmailUserError] = useState(false);
	const [cityUser, setCityUser] = useState("");
	const [telegramLinkUser, setTelegramLinkUser] = useState("");
	const [telegramLinkUserError, setTelegramLinkUserError] = useState(false);
	const [birthdayUser, setBirthdayUser] = useState("");

	const [token, setToken] = useLocalStorage({
		initialValue: "",
		key: "token",
	});

	const [id_user, setIdUser] = useLocalStorage({
		initialValue: -1,
		key: "id_user",
	});

	const [profile_id, setProfileId] = useLocalStorage({
		initialValue: -1,
		key: "profile_id",
	});

	const navigate = useNavigate();

	return (
		<div className={styles.wrapper__contacts}>
			<div className={styles.wrapper__userPath}>
				{<UserPath Stage={StageType.First} />}
			</div>
			<div className={styles.wrapper__contacts__maincontent}>
				<div className={styles.contacts__mainHeader}>
					<h1 className={styles.mainHeader}>Добро пожаловать!</h1>
				</div>
				<div className={styles.contacts__header__h2}>
					<h2 className={styles.header__h2}>
						Помогите нам узнать вас лучше
					</h2>
				</div>
				<form
					onSubmit={async (event) => {
						event.preventDefault();

						const data = {
							city: cityUser,
							findstatus: true,
							name: nameUser,
							photo: "1",
							surname: surnameUser,
							telegram: telegramLinkUser,
							birthday: birthdayUser,
							country: "string",
							id: id_user,
						};

						if (data.city.length == 0) {
							data.city = "none";
						}

						if (data.name.length == 0) {
							setNameUserError(true);
							return;
						}
						if (data.surname.length == 0) {
							setSurnameUserError(true);
							return;
						}
						if (data.telegram.length == 0) {
							setTelegramLinkUserError(true);
							return;
						}

						const response = await fetch(
							"http://localhost:8000/api/profile/create_profile",
							{
								method: "POST",
								body: JSON.stringify(data),
								headers: {
									"Content-Type": "application/json",
									Authorization: `Bearer ${token}`,
								},
								credentials: "include",
							},
						);
						if (!response.ok) {
							return;
						}

						const dataFromResponse = await response.json();

						setProfileId(dataFromResponse.profile_id);

						setTimeout(() => {
							navigate("/auth_profile/activity", {
								replace: true,
							});
						}, 1);
					}}
					action="http://localhost:8000/api/profile/create_profile"
					method="post"
					className={styles.contacts__form}
				>
					<div className={styles.contacts}>
						<div className={styles.wrapper__inputs}>
							<div className={styles.wrapper__input__lfs}>
								{
									<TextField
										id={"name"}
										textLabel={"Имя"}
										typeInput={"text"}
										inputData={nameUser}
										setInput={setNameUser}
										setErrorInput={setNameUserError}
										error={nameUserError}
										location={LocationInputField.Profile}
										typeInputOnProfile={
											TypeInputOnProfile.Double
										}
									/>
								}
							</div>
							<div className={styles.wrapper__input__rfs}>
								{
									<TextField
										id={"surname"}
										textLabel={"Фамилия"}
										typeInput={"text"}
										inputData={surnameUser}
										setInput={setSurNameUser}
										setErrorInput={setSurnameUserError}
										error={surnameUserError}
										location={LocationInputField.Profile}
										typeInputOnProfile={
											TypeInputOnProfile.Double
										}
									/>
								}
							</div>
						</div>
						<div className={styles.wrapper__inputs}>
							{
								<TextField
									id={"email"}
									textLabel={"Email"}
									typeInput={"email"}
									inputData={emailUser}
									setInput={setEmailUser}
									setErrorInput={setEmailUserError}
									error={emailUserError}
									location={LocationInputField.Profile}
									typeInputOnProfile={
										TypeInputOnProfile.Single
									}
								/>
							}
						</div>
						<div className={styles.wrapper__inputs}>
							<div className={styles.wrapper__input__lfs}>
								{
									<TextField
										id={"tgLink"}
										textLabel={"Ссылка на телеграм"}
										typeInput={"url"}
										inputData={telegramLinkUser}
										setInput={setTelegramLinkUser}
										setErrorInput={setTelegramLinkUserError}
										error={telegramLinkUserError}
										location={LocationInputField.Profile}
										typeInputOnProfile={
											TypeInputOnProfile.Double
										}
									/>
								}
							</div>
							<div className={styles.wrapper__input__rfs}>
								{
									<TextField
										id={"birthday"}
										textLabel={"Дата рождения"}
										typeInput={"date"}
										inputData={birthdayUser}
										setInput={setBirthdayUser}
										location={LocationInputField.Profile}
										typeInputOnProfile={
											TypeInputOnProfile.Double
										}
									/>
								}
							</div>
						</div>
						<div className={styles.wrapper__inputs}>
							{
								<TextField
									id={"city"}
									textLabel={"Город"}
									typeInput={"text"}
									inputData={cityUser}
									setInput={setCityUser}
									location={LocationInputField.Profile}
									typeInputOnProfile={
										TypeInputOnProfile.Single
									}
								/>
							}
						</div>
					</div>
					<div className={styles.profile__footer}>
						<div className={styles.wrapper__buttonNext}>
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
