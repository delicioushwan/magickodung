import { useForm } from "react-hook-form";

export default function Login() {
  const { register, handleSubmit } = useForm();
  const onSubmit = (data) => console.log(data);
  return (
    <div>
      <p>로그인</p>
      <form onSubmit={handleSubmit(onSubmit)}>
        아이디
        <input {...register("account", { required: true })} />
        비번
        <input type="password" {...register("pwd", { required: true })} />
        <input type="submit" />
      </form>
    </div>
  );
}
