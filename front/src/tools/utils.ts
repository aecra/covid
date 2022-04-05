let buildQueryString = (params: any) => {
  let queryString = '';
  for (let key in params) {
    if (params.hasOwnProperty(key)) {
      queryString += encodeURIComponent(key) + '=' + encodeURIComponent(params[key]) + '&';
    }
  }
  return queryString;
};

let to = async (promise: Promise<any>) => {
  try {
    const data = await promise;
    return [null, data];
  } catch (err) {
    return [err];
  }
};

export { buildQueryString, to };
