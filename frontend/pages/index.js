import Link from "next/link";

export default function Home() {
  return (
    <div>
      <div>
        <Link href="/quesions/add">질문하기</Link>
      </div>
      <div>
        <Link href="/questions">골라주기</Link>
      </div>
      <div>
        <Link href="/login">로그인</Link>
      </div>
      <div>
        <Link href="/signup">회원되기</Link>
      </div>
    </div>
  );
}
