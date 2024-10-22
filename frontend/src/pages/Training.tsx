import React, { useState, useEffect } from 'react';
import './Training.css';

const Training = () => {
  const [problem, setProblem] = useState<{ shape1: string; shape2: string; option: number; withQ: boolean } | null>(null);
  const [correctCount, setCorrectCount] = useState<number | null>(null);
  const [wrongCount, setWrongCount] = useState<number | null>(null);

  const fetchProblem = async () => {
    try {
      const response = await fetch('http://localhost:3000/problem');
      const data = await response.json();
      setProblem(data);
    } catch (error) {
      console.error('問題の取得に失敗しました:', error);
    }
  };

  const fetchSummary = async () => {
    try {
      const response = await fetch('http://localhost:3000/summary');
      if (response.ok) {
        const result = await response.json();
        setCorrectCount(result.correctCount);
        setWrongCount(result.wrongCount);
      }
    } catch (error) {
      console.error('集計結果の取得に失敗しました:', error);
    }
  };

  const sendAnswer = async (answer: string) => {
    try {
      const response = await fetch('http://localhost:3000/answer', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ answer }),
      });

      if (response.ok) {
        fetchProblem();
        fetchSummary();
      } else {
        alert('サーバーに送信中にエラーが発生しました。');
      }
    } catch (error) {
      console.error('答えの送信に失敗しました:', error);
    }
  };

  useEffect(() => {
    fetchProblem();
  }, []);

  return (
    <div className="training-container">
      <h2 className="training-title">
        Training
        {correctCount !== null && wrongCount !== null && (
          <span className="result-summary">
            (正解: {correctCount}, 不正解: {wrongCount})
          </span>
        )}
      </h2>
      {problem ? (
        <div>
          <p className="problem-display">
            問題: {problem.withQ ? 'Q' : ''}{problem.shape1} {problem.shape2}
          </p>
          <button className="answer-button" onClick={() => sendAnswer('q')}>
            !
          </button>
          <button className="answer-button" onClick={() => sendAnswer('w')}>
            {problem.option}
          </button>
          <button className="answer-button" onClick={() => sendAnswer('e')}>
            E
          </button>
        </div>
      ) : (
        <p>問題を読み込み中...</p>
      )}
    </div>
  );
};

export default Training;
