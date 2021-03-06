import { useState } from "react";
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
    </div>
  );
}

async function onSubmitSignup({ account, pwd }, setLoading) {
  await Axios.post("/users/signup", { account, pwd });
  setLoading(false);
}
