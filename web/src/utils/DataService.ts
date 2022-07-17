import OAuthService from './OAuthService';
import { to } from './utils';

const DataService = {
  getUser: async () => {
    const [err, response] = await to(OAuthService.Request().post('/getUser'));
    if (err != null) {
      return [Error('Request error'), {}];
    }
    const data = response.data;
    return [
      null,
      data,
    ];
  },
  updateUser: async (data: {
    name: string;
    email: string;
    state: boolean;
    position: string;
    eaisess: string;
    uukey: string;
    home: string;
  }) => {
    const [err] = await to(OAuthService.Request().post('/updateUser', data));
    if (err != null) {
      return [Error('Request error')];
    }
    return [null];
  },
};

export default DataService;
