const HOST = 'http://127.0.0.1:8080';

const LOGIN = HOST + '/api/admin/auth/login';
const ME = HOST + '/api/admin/auth/me';

const POSTS = HOST + '/api/admin/posts';

const CATEGORIES = HOST + '/api/admin/categories-all';
const CATEGORIES_LIST = HOST + '/api/admin/categories';

const PAGES = HOST + '/api/admin/pages';

const DASHBOARDS = HOST + '/api/admin/dashboards';

export default {
  LOGIN,
  ME,
  POSTS,
  CATEGORIES,
  CATEGORIES_LIST,
  PAGES,
  DASHBOARDS,
};
