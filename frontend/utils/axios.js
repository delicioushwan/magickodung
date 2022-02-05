import Axios from "axios";
console.log(process.env.NEXT_PUBLIC_BASE_URL);
const instance = Axios.create({
  withCredentials: true,
  baseURL: process.env.NEXT_PUBLIC_BASE_URL,
});

export default instance;
