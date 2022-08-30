import auth from './auth';
import {to} from './utils';

const DataService = {
  getUser: async () => {
    const [err, response] = await to(auth.GetRequest().get('/api/user/profile'));
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
    username: string;
    email: string;
    state: boolean;
    position: string;
    eaisess: string;
    uukey: string;
    home: string;
  }) => {
    const [err] = await to(auth.GetRequest().put('/api/user/profile', data));
    if (err != null) {
      return [Error('Request error')];
    }
    return [null];
  },
  report: async () => {
    const [err] = await to(auth.GetRequest().post('/api/user/reportTest'));
    if (err != null) {
      return [Error('Request error')];
    }
    return [null];
  },
  getRecords: async () => {
    const [err, response] = await to(auth.GetRequest().get('/api/user/records'));
    if (err != null) {
      return [Error('Request error'), {}];
    }
    const data = response.data;
    return [
      null,
      data,
    ];
  }
};

export default DataService;
