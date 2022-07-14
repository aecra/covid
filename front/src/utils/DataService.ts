import OAuthService from './OAuthService';
import { to } from './utils';

const DataService = {
  getUser: async () => {
    const [err, response] = await to(OAuthService.Request().post('/User'));
    if (err != null) {
      return [Error('Request error'), {}];
    }
    const data = response.data;
    return [
      null,
      {
        name: data.name,
        email: data.email,
        state: data.state === 'on',
        position: data.position,
        eaisess: data.eaisess,
        uukey: data.uukey,
        home: data.home,
      },
    ];
  },
  updateUser: async (data: {
    name: string;
    email: string;
    state: string;
    position: string;
    eaisess: string;
    uukey: string;
    home: string;
  }) => {
    const [err] = await to(OAuthService.Request().post('/UpdateUser', data));
    if (err != null) {
      return [Error('Request error')];
    }
    return [null];
  },
};

export default DataService;
