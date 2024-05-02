import { useEffect, useState } from "react";

type propsUseLocalStorage = {
	initialValue: any;
	key: string;
};

function useLocalStorage(props: propsUseLocalStorage) {
	const { initialValue, key } = props;

	const getValue = () => {
		const storage = localStorage.getItem(key);

		if (storage) {
			return JSON.parse(storage);
		}

		return initialValue;
	};

	const [value, setValue] = useState(getValue);

	useEffect(() => {
		localStorage.setItem(key, JSON.stringify(value));
	}, [value]);

	return [value, setValue];
}

export { useLocalStorage };
