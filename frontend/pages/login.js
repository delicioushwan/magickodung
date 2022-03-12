import { useForm } from "react-hook-form";
import { useState } from "react";
import { useRouter } from "next/router";

import Axios from "../utils/axios";

export default function Login() {
  const { register, handleSubmit } = useForm();
  const [isLoading, setLoading] = useState(false);
  const router = useRouter();
  const callback = () => {
    setLoading(false);
    window.alert("로그인 되었습니다.");
    router.replace("/");
  };
  const onSubmit = async (data) => {
    setLoading(true);
    await onSubmitLogin(data, callback);
  };
  return (
    <div>
      <p>로그인</p>
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

async function onSubmitLogin({ account, pwd }, callback) {
  try {
    await Axios.post("/users/login", { account, pwd });
    callback();
  } catch (e) {
    console.log(e);
  }
}
