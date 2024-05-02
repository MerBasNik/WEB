import React from "react";
import { Route, Routes } from "react-router-dom";
import { MainPage } from "../mainPage/mainPage";
import { RegistrationScreen } from "../registerPage/register";
import { ReEntryScreen } from "../registerPage/re-entry";
import { Profile } from "../profilePage/profile";

export const Main = () => {
	return (
		<Routes>
			<Route path="/" element={<RegistrationScreen />} />
			<Route path="/reEnt" element={<ReEntryScreen />} />
			<Route path="/main" element={<MainPage />}>
				<Route index element={<></>} />
				<Route path="/main/chats" element={<></>} />
				<Route path="/main/profile" element={<Profile />} />
			</Route>
		</Routes>
	);
};
