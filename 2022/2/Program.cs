namespace Elves
{
    class Program
    {
        enum Move
        {
            Rock,
            Paper,
            Scissors
        }

		enum Result
		{
			PlayerWin,
			Draw,
			OpponentWin
		}

        static void Main()
        {
            string[] lines = File.ReadAllLines("input.txt");

			uint totalPlayerScore = 0;

            for (uint i = 0; i < lines.Length; i++)
            {
                (Move linePlayerMove, Move lineOpponentMove, Result lineDesiredResult) = Parse(lines[i]);

				linePlayerMove = CalculateMove(lineOpponentMove, lineDesiredResult);

				Result lineMatchResult = CalculateMatchResult(linePlayerMove, lineOpponentMove);

				uint linePlayerScore = CalculatePlayerScore(linePlayerMove, lineMatchResult);

				totalPlayerScore += linePlayerScore;
            }

			Console.WriteLine("Player total score: " + totalPlayerScore);
        }

        static (Move player, Move opponent, Result desired) Parse(string line)
        {
            Move player = Move.Rock;
			Move opponent = Move.Rock;
			Result desired = Result.PlayerWin;

			if(line[0].Equals('A')) opponent = Move.Rock;
			else if (line[0].Equals('B')) opponent = Move.Paper;
			else if (line[0].Equals('C')) opponent = Move.Scissors;

			if(line[2].Equals('X'))
			{
				player = Move.Rock;
				desired = Result.OpponentWin;
			}
			else if (line[2].Equals('Y'))
			{
				player = Move.Paper;
				desired = Result.Draw;
			}
			else if (line[2].Equals('Z'))
			{
				player = Move.Scissors;
				desired = Result.PlayerWin;
			}

			return (player, opponent, desired);
        }

		static Result CalculateMatchResult(Move player, Move opponent)
		{
			// Player wins by default
			Result r = Result.PlayerWin;

			// Draw if the same
			if (player.Equals(opponent)) r = Result.Draw;
			else if(player == Move.Rock && opponent == Move.Paper
			|| player == Move.Paper && opponent == Move.Scissors
			|| player == Move.Scissors && opponent == Move.Rock)
			{
				r = Result.OpponentWin;
			}

			return r;
		}

		static Move CalculateMove(Move opponent, Result desired)
		{
			// Rock by default
			Move m = Move.Rock;

			if (desired == Result.Draw) m = opponent;
			else if(desired == Result.PlayerWin)
			{
				if(opponent == Move.Rock) m = Move.Paper;
				else if (opponent == Move.Paper) m = Move.Scissors;
			}
			else if(desired == Result.OpponentWin)
			{
				if(opponent == Move.Scissors) m = Move.Paper;
				else if (opponent == Move.Rock) m = Move.Scissors;
			}

			return m;
		}

		static uint CalculatePlayerScore(Move m, Result r)
		{
			uint score = 0;

			if (m == Move.Rock) score += 1;
			else if (m == Move.Paper) score += 2;
			else if (m == Move.Scissors) score += 3;

			if (r == Result.PlayerWin) score += 6;
			else if (r == Result.Draw) score += 3;
			else if (r == Result.OpponentWin) score += 0;

			return score;
		}
    }
}