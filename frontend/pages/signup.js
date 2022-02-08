import { useState, useffect, useEffect } from "react";
import { useForm } from "react-hook-form";
import Axios from "../utils/axios";

export default function Login() {
  const { register, handleSubmit } = useForm();
  const [isLoading, setLoading] = useState(false);

  const onSubmit = async (data) => {
    setLoading(true);
    await onSubmitSignup(data, setLoading);
  };

  return (
    <div>
      <p>회원가입</p>
      <form onSubmit={handleSubmit(onSubmit)}>
        아이디
        <input {...register("account", { required: true })} />
        비번
        <input type="password" {...register("pwd", { required: true })} />
        <input type="submit" disabled={isLoading} />
      </form>
      <button onClick={onPost}>post</button>
      <button onClick={onGet}>get</button>
    </div>
  );
}

async function onSubmitSignup({ account, pwd }, setLoading) {
  const response = await Axios.post("/users/signup", { account, pwd });
  console.log(response);
  setLoading(false);
}

async function onPost({ account, pwd }, setLoading) {
  const response = await Axios.post("/test");
  console.log(response);
}

async function onGet({ account, pwd }, setLoading) {
  const response = await Axios.get("/test");
  console.log(response);
}
