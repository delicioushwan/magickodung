import React, { useEffect } from 'react';
import styled from 'styled-components';
import Link from 'next/link';

export default function Layout({ children, isLogin }) {
  useEffect(() => {
    window.scrollTo({ top: 0, left: 0, behavior: 'auto' });
  }, []);

  return (
    <Wrapper>
      <div>
        <NaviWrapper>
          <Navi>
            <Link href="/questions/add">질문하기</Link>
          </Navi>
          <Navi>
            <Link href="/questions">골라주기</Link>
          </Navi>
          {isLogin ? (
            <Navi>
              <Link href="/my-page">My Page</Link>
            </Navi>
          ) : (
            <>
              <Navi>
                <Link href="/login">로그인</Link>
              </Navi>
              <Navi>
                <Link href="/signup">회원되기</Link>
              </Navi>
            </>
          )}
        </NaviWrapper>
        <div>{children}</div>
      </div>
    </Wrapper>
  );
}

const Wrapper = styled.div`
  display: flex;
  justify-content: center;
  padding: 25px 10px;
  font-family: Arial;
`;

const NaviWrapper = styled.div`
  display: flex;
  flex-wrap: wrap;
  margin-bottom: 20px;
  justify-content: center;
`;

const Navi = styled.div`
  & + & {
    margin-left: 10px;
  }

  a {
    text-decoration: none;
    outline: none;
    color: black;
  }

  a:hover,
  a:active {
    text-decoration: none;
    color: #fff;
    background-color: #f59000;
  }
`;
