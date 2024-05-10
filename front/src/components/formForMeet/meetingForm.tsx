import React, { useState } from "react";
import stylesForm from "./meetingForm.module.css";
import stylesContacts from "../../views/profilePage/contacts/contactsPage.module.css";
import {
	LocationInputField,
	TextField,
	TypeInputOnProfile,
} from "../input/input";
import { Button, ButtonType } from "../button/button";
import { useLocalStorage } from "../../hooks/useLocalStorage";

type MeetingFormProps = {
	active: boolean;
	setActive: (active: boolean) => void;
};

export const MeetingForm = (props: MeetingFormProps) => {
	const { active, setActive } = props;

	const [meetingName, setMeetingName] = useState("");

	const [dayStart, setDayStart] = useState("");
	const [errorDayStart, setErrorDayStart] = useState(false);
	const [dayEnd, setDayEnd] = useState("");
	const [errorDayEnd, setErrorDayEnd] = useState(false);

	const [timeStart, setTimeStart] = useState("");
	const [errorTimeStart, setErrorTimeStart] = useState(false);
	const [timeEnd, setTimeEnd] = useState("");
	const [errorTimeEnd, setErrorTimeEnd] = useState(false);

	const [checkboxAnswer, setCheckboxAnswer] = useState(false);

	let classesForForm =
		stylesForm.meetingForm + " " + stylesForm.meetingForm__notActive;
	if (active) {
		classesForForm = stylesForm.meetingForm;
	} else {
		classesForForm =
			stylesForm.meetingForm + " " + stylesForm.meetingForm__notActive;
	}

	const [token] = useLocalStorage({
		initialValue: "",
		key: "token",
	});

	const [id_user] = useLocalStorage({
		initialValue: "",
		key: "id_user",
	});

	const dataTime = new Date();

	const hourMinute = `${dataTime.getHours()}:${dataTime.getMinutes()}`;

	return (
		<div className={classesForForm}>
			<form
				onSubmit={async (event) => {
					event.preventDefault();
					const data = {
						endday: dayEnd,
						endtime: timeEnd,
						startday: dayStart,
						starttime: timeStart,
					};

					const arrDayStart = dayStart.split("-");
					const yearMonthDayStart = arrDayStart.map((str) =>
						parseInt(str),
					);
					const dayStartData = new Date(
						yearMonthDayStart[0],
						yearMonthDayStart[1],
						yearMonthDayStart[2],
					);

					const arrDayEnd = dayEnd.split("-");
					const yearMonthDayEnd = arrDayEnd.map((str) =>
						parseInt(str),
					);

					const dayEndData = new Date(
						yearMonthDayEnd[0],
						yearMonthDayEnd[1],
						yearMonthDayEnd[2],
					);

					if (
						timeStart < hourMinute &&
						dayStartData.getTime() === dataTime.getTime()
					) {
						setErrorTimeStart(true);
						return;
					}

					if (timeEnd <= timeStart) {
						setErrorTimeEnd(true);
						return;
					}

					if (dayStartData.getTime() < dataTime.getTime()) {
						setErrorTimeStart(true);
						return;
					}

					if (dayEndData.getTime() < dayStartData.getTime()) {
						setErrorDayEnd(true);
						return;
					}

					let url = "";
					if (!checkboxAnswer) {
						url =
							"http://localhost:8000/api/chats/find_chats_users";
					} else {
						url =
							"http://localhost:8000/api/chats/find_chats_users_by_hobby";
					}

					const response = await fetch(url, {
						method: "POST",
						body: JSON.stringify(data),
						headers: {
							"Content-Type": "application/json",
							Authorization: `Bearer ${token}`,
						},
						credentials: "include",
					});

					if (response.ok) {
						const dataFromResponse = await response.json();

						const dataForChat = {
							title: "ЫЫЫЫЫЫ",
							usersId: [
								id_user,
								dataFromResponse.finded_user_id_for_chat,
							],
						};

						const responceCreateChat = await fetch(
							"http://localhost:8000/api/chats/create_chat",
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

						// if (responceCreateChat.ok) {
						// }
					}

					setTimeout(() => {
						setActive(!active);
					}, 1);
				}}
				action="http://localhost:8000/api/profile/create_profile"
				method="post"
			>
				<div className={stylesForm.wrapper__meetingForm__header}>
					<h1 className={stylesForm.meetingForm__header}>
						Создание встречу
					</h1>
				</div>
				<div className={stylesContacts.wrapper__inputs}>
					{
						<TextField
							id={"MeetingName"}
							textLabel={"Название встречи"}
							typeInput={"text"}
							inputData={meetingName}
							setInput={setMeetingName}
							location={LocationInputField.Profile}
							typeInputOnProfile={TypeInputOnProfile.Single}
						/>
					}
				</div>
				<div className={stylesContacts.wrapper__inputs}>
					<div className={stylesContacts.wrapper__input__lfs}>
						{
							<TextField
								id={"dayStart"}
								textLabel={"День начало"}
								typeInput={"date"}
								inputData={dayStart}
								setInput={setDayStart}
								error={errorDayStart}
								setErrorInput={setErrorDayStart}
								location={LocationInputField.Profile}
								typeInputOnProfile={TypeInputOnProfile.Double}
							/>
						}
					</div>
					<div className={stylesContacts.wrapper__input__rfs}>
						{
							<TextField
								id={"dayEnd"}
								textLabel={"День конец"}
								typeInput={"date"}
								inputData={dayEnd}
								setInput={setDayEnd}
								error={errorDayEnd}
								setErrorInput={setErrorDayEnd}
								location={LocationInputField.Profile}
								typeInputOnProfile={TypeInputOnProfile.Double}
							/>
						}
					</div>
				</div>
				<div className={stylesContacts.wrapper__inputs}>
					<div className={stylesContacts.wrapper__input__lfs}>
						{
							<TextField
								id={"timeStart"}
								textLabel={"Время начало"}
								typeInput={"time"}
								inputData={timeStart}
								setInput={setTimeStart}
								error={errorTimeStart}
								setErrorInput={setErrorTimeStart}
								location={LocationInputField.Profile}
								typeInputOnProfile={TypeInputOnProfile.Double}
							/>
						}
					</div>
					<div className={stylesContacts.wrapper__input__rfs}>
						{
							<TextField
								id={"timeEnd"}
								textLabel={"Время конец"}
								typeInput={"time"}
								inputData={timeEnd}
								setInput={setTimeEnd}
								error={errorTimeEnd}
								setErrorInput={setErrorTimeEnd}
								location={LocationInputField.Profile}
								typeInputOnProfile={TypeInputOnProfile.Double}
							/>
						}
					</div>
				</div>
				<div className={stylesForm.meetingForm__addParams}>
					<div className={stylesForm.addParams__input}>
						<input
							id="AnswerYes"
							type="checkbox"
							className={stylesForm.input__checkbox}
							onChange={(event) => {
								event.preventDefault();
								setCheckboxAnswer(!checkboxAnswer);
							}}
						/>
					</div>
					<div className={stylesForm.addParams__label}>
						<label
							htmlFor="AnswerYes"
							className={stylesForm.input__label}
						>
							Встреча по интересам
						</label>
					</div>
				</div>
				{/* <div className={stylesForm.meetingForm__result}></div> */}
				<div className={stylesForm.meetingForm__buttons}>
					<div className={stylesForm.wrapper__button__cancel}>
						<Button
							title={"Закрыть"}
							id={"cancel"}
							type={ButtonType.Text}
							typeButton={"button"}
							onClick={() => {
								setActive(false);
							}}
						/>
					</div>
					<div className={stylesForm.wrapper__button__createEvent}>
						<Button
							title={"Создать"}
							id={"create"}
							type={ButtonType.Text}
							typeButton={"submit"}
						/>
					</div>
				</div>
			</form>
		</div>
	);
};
