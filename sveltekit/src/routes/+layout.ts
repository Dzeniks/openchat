export const trailingSlash = 'always';
export const ssr = false;

export const load = ({ url }) => {
	const { pathname } = url;

	return {
		pathname
	};
    };
