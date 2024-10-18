let problem = null;
let score = { correct: 0, incorrect: 0 };
let questionCount = 0;

const shapeValues = {
    "〇": 0,
    "△": 3,
    "□": 4,
    "☆": 5
};

async function loadProblem() {
    try {
        const response = await fetch('http://localhost:8080/problem');
        problem = await response.json();
        displayProblem();
    } catch (error) {
        console.error('Error loading problem:', error);
    }
}

function displayProblem() {
    const problemElement = document.getElementById('problem');
    problemElement.textContent = `${problem.withQ ? 'Q' : ''}${problem.shape1} ${problem.shape2}`;
    
    const optionButton = document.getElementById('optionButton');
    optionButton.textContent = problem.option;
    optionButton.setAttribute('aria-label', `オプション ${problem.option}`);
}
function handleAnswer(answer) {
    if (!problem) return;

    let isCorrect = false;
    const correctSum = shapeValues[problem.shape1] + shapeValues[problem.shape2];

    if (problem.withQ && answer === 'q') {
        isCorrect = true;
    } else if (!problem.withQ && answer === 'w' && problem.option === correctSum) {
        isCorrect = true;
    } else if (!problem.withQ && answer === 'e' && problem.option !== correctSum) {
        isCorrect = true;
    }

    if (isCorrect) {
        score.correct++;
    } else {
        score.incorrect++;
    }

    questionCount++;
    updateScore();

    if (questionCount < 10) {
        loadProblem();
    } else {
        endTest();
    }
}

function updateScore() {
    const scoreElement = document.getElementById('score');
    scoreElement.textContent = `正解: ${score.correct}, 不正解: ${score.incorrect}`;
}

function endTest() {
    const problemElement = document.getElementById('problem');
    problemElement.textContent = '練習終了';
    
    const buttonsElement = document.getElementById('buttons');
    buttonsElement.style.display = 'none';
}

loadProblem();