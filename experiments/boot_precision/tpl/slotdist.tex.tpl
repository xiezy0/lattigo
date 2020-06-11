\begin{tikzpicture}
    \begin{axis}[
        title={},
        width=0.5\textwidth,
        height=0.5\textwidth,
        xlabel={Precision (bits)},
        ylabel={$\Pr[\mathsf{prec}(\mathsf{slot})<x]$},
        xmin=15, xmax=32,
        ymin=0, ymax=1,
        legend pos=north west,
        ymajorgrids=true,
        grid style=dashed],
    ]
    
    \addplot[only marks, color=blue, mark size=.5] coordinates {
% Real
{{.DataReal}}
    };
    \addplot[only marks, color=red, mark size=.5] coordinates {
% Imag
{{.DataImag}}
    };
    \legend{Real,Imag}
        
    \end{axis}
\end{tikzpicture}