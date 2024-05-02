import React from "react";
import styles from "./button.module.css";
import { Link } from "react-router-dom";

export enum ButtonType {
	Icon,
	Text,
}

type ButtonProps = {
	onClick?: () => void;
	title?: string;
	id: string;
	type: ButtonType;
	typeButton?: "button" | "submit" | "reset";
	icon?: JSX.Element | undefined;
};

export const Button = (props: ButtonProps) => {
	const { onClick, id, icon, type, title, typeButton } = props;
	if (type == ButtonType.Icon) {
		if (icon) {
			if (onClick) {
				return (
					<div key={id} className={styles.buttonWrapper}>
						<button
							key={id}
							onClick={(
								eventt: React.MouseEvent<HTMLButtonElement>,
							) => {
								eventt.preventDefault();
								onClick();
							}}
							className={styles.buttonIcon}
						>
							{icon}
						</button>
					</div>
				);
			} else {
				return (
					<div key={id} className={styles.buttonWrapper}>
						<button key={id} className={styles.buttonIcon}>
							{icon}
						</button>
					</div>
				);
			}
		} else {
			return <></>;
		}
	}
	if (type == ButtonType.Text) {
		if (onClick) {
			return (
				<div key={id} className={styles.buttonWrapper}>
					<button
						key={id}
						onClick={(
							eventt: React.MouseEvent<HTMLButtonElement>,
						) => {
							eventt.preventDefault();
							onClick();
						}}
						className={styles.buttonText}
						type={typeButton}
					>
						{title}
					</button>
				</div>
			);
		} else {
			return (
				<div key={id} className={styles.buttonWrapper}>
					<button
						key={id}
						className={styles.buttonText}
						type={typeButton}
					>
						{title}
					</button>
				</div>
			);
		}
	}
	return <></>;
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

export enum LocationOnPages {
	Authorization,
	HomePage,
}

type ButtonTextLinkProps = {
	id: string;
	link: string;
	title: string;
	location: LocationOnPages;
};

export const ButtonTextLink = (props: ButtonTextLinkProps) => {
	const { id, link, title, location } = props;

	return (
		<div key={id} className={styles.buttonWrapper}>
			<Link key={id} to={link} className={styles.buttonText}>
				{title}
			</Link>
		</div>
	);
};
