import React from "react"
import styles from './button.module.css'
import stylesRegister from "../../views/registerPage/register.module.css"
import { Link, NavLink } from "react-router-dom"
import { useState } from 'react'

export enum ButtonType {
	Icon,
	Text,
}

type ButtonProps = {
	onClick: () => void;
	title?: string;
	id: string;
	type: ButtonType;
	icon?: JSX.Element | undefined;
	link?: string;
};

export const Button = (props: ButtonProps) => {
	const { onClick, id, icon, type, title, link } = props;
	if (type == ButtonType.Icon) {
		if (icon) {
			return (
				<div key={id} className={styles.buttonWrapper}>
					<button
						key={id}
						onClick={onClick}
						className={styles.buttonIcon}
					>
						{icon}
					</button>
				</div>
			);
		}
	}
	if (type == ButtonType.Text) {
		return (
			<div key={id} className={styles.buttonWrapper}>
				<button key={id} onClick={onClick} className={styles.button}>
					{title}
				</button>
			</div>
		);
	}
	return <div />;
};

type ButtonIconLinkProps = {
	id: string;
	link: string;
	icon: JSX.Element;
};

export const ButtonIconLink = (props: ButtonIconLinkProps) => {
	const { id, link, icon } = props;
	return (
		<div key={id} className={styles.buttonWrapper}>
			<Link key={id} to={link} className={styles.buttonIcon}>
				{icon}
			</Link>
		</div>
	);
};

type ButtonTextLinkProps = {
	id: string;
	link: string;
	title: string;
};

export const ButtonTextLink = (props: ButtonTextLinkProps) => {
	const { id, link, title } = props;
	return (
		<div key={id} className={styles.buttonWrapper}>
			<Link key={id} to={link} className={stylesRegister.input__submit}>
				{title}
			</Link>
		</div>
	);
};
