export const getData = async (url: string) => {
	const response = await fetch(url);

	if (!response.ok) {
		throw new Error(
			`Ошибка по адресу ${url}, статус ошибки ${response.status}`,
		);
	}
	return await response.json();
};

export const sendData = async (url: string, data: any) => {
	const response = await fetch(url, {
		method: "POST",
		body: JSON.stringify(data),
	});

	if (!response.ok) {
		throw new Error(
			`Ошибка по адресу ${url}, статус ошибки ${response.status}`,
		);
	}
	return await response.json();
};
