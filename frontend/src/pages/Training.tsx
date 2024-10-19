import React, { useState, useEffect } from 'react';

const Training = () => {
  const [problem, setProblem] = useState<{ shape1: string; shape2: string; correctSum: number; withQ: boolean } | null>(null);

  // 問題を取得する関数
  const fetchProblem = async () => {
    try {
      const response = await fetch('http://localhost:3000/problem');
      const data = await response.json();
      setProblem(data);
    } catch (error) {
      console.error('問題の取得に失敗しました:', error);
    }
  };

  // 初回レンダリング時に問題を取得
  useEffect(() => {
    fetchProblem();
  }, []);

  return (
    <div>
      <h2>Training</h2>
      {problem ? (
        <div>
          <p>問題: {problem.withQ ? 'Q' : ''}{problem.shape1} {problem.shape2}</p>
          <p>合計: {problem.correctSum}</p>
          <button onClick={() => alert('qを選択')}>q: !</button>
          <button onClick={() => alert('wを選択')}>w: {problem.correctSum}</button>
          <button onClick={() => alert('eを選択')}>e: E</button>
        </div>
      ) : (
        <p>問題を読み込み中...</p>
      )}
    </div>
  );
};

export default Training;
