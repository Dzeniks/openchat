export async function load({ cookies }) {
    const accessToken = cookies.get('accessToken');
    const refreshToken = cookies.get('refreshToken');

    console.log('accessToken:', accessToken);
    console.log('refreshToken:', refreshToken);
    return {
        accessToken: accessToken,
        refreshToken: refreshToken
      };
    }