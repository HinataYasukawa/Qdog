import React from 'react';
import './Home.css'; 

const Home = () => {
  return (
    <div className="home-container">
      <h2 className="home-title">Home</h2>
      <p className="home-description">
        本アプリケーションは適性検査Qdogのうちの一つである図形問題の練習を行うものです。<br />
        二つの図形が表示されるので、その角の数の合計を回答するものです。
      </p>
      <p className="home-instructions">
        選択肢は「!」「数字」「E」の三種類です：
        <ul>
          <li>選択肢に正解の数字が書かれている場合：「数字」を選ぶ</li>
          <li>選択肢に正解の数字が書かれていない場合：「E」を選ぶ</li>
          <li>図形の前にQが表示されている場合：「!」を選ぶ</li>
        </ul>
        10問ごとに正解数が表示されます。
      </p>
      <div className="home-footer">
        <a 
          href="https://github.com/HinataYasukawa/Qdog" 
          target="_blank" 
          rel="noopener noreferrer"
          className="github-link"
        >
          GitHub Repository
        </a>
      </div>
    </div>
  );
};

export default Home;
