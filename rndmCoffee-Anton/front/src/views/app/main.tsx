import React from "react";
import { Route, Routes } from "react-router-dom";
import { RegistrationScreen } from "../registerPage/register";
import { ReEntryScreen } from "../registerPage/re-entry";
import { MainProfile } from "../profilePage/mainProfile/mainProfilePage";
import { UserContacts } from "../profilePage/contacts/contactsPage";
import { WelcomePage } from "../profilePage/welcome/welcomePage";
import { MainPage } from "../mainPage/mainPage";
import { ActivityPage } from "../profilePage/activity/activityPage";
import { HomePage } from "../mainPage/homePage/homePage";

export const Main = () => {
	return (
		<Routes>
			<Route path="/" element={<ReEntryScreen />} />
			<Route path="/reg" element={<RegistrationScreen />} />
			<Route path="/auth_profile" element={<MainProfile />}>
				<Route
					path="/auth_profile/contacts"
					element={<UserContacts />}
				/>
				<Route
					path="/auth_profile/activity"
					element={<ActivityPage />}
				/>
				<Route path="/auth_profile/welcome" element={<WelcomePage />} />
			</Route>
			<Route path="/main" element={<MainPage />}>
				<Route path="/main/home" element={<HomePage />} />
				<Route path="/main/profile" element={<></>} />
			</Route>
		</Routes>
	);
};
