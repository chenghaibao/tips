
<?php
$tree = [

    "A" => ["B", "C","D"],

    "B" => ["A","E", "F"],

    "C" => ["G","E","A"],

    "D" => ["A"],

    "F" => ["B","E"],

    "E" => ["B","C", "F","H"],

    "G" =>["C"],

    "H" =>["E"],

];

bfs($tree, "A");

echo PHP_EOL;

dfs($tree, "A");

function bfs($tree, $node) {

    $queue = [$node];

    $usedQueue = [

        $node =>true

    ];

    while (!empty($queue)) {

        echo $node = array_shift($queue);

        if (!isset($tree[$node])) continue;

        foreach ($tree[$node] as $childNode) {

            if (isset($usedQueue[$childNode])) {

                continue;

            }

            array_push($queue, $childNode);

            $usedQueue[$childNode] = true;

        }

    }

}

function dfs($tree, $node) {

    $queue = [$node];

    $usedQueue = [

        $node =>true

    ];

    while (!empty($queue)) {

        echo $node = array_pop($queue);

        if (!isset($tree[$node])) continue;

        foreach ($tree[$node] as $childNode) {

            if (isset($usedQueue[$childNode])) {

                continue;

            }

            array_push($queue, $childNode);

            $usedQueue[$childNode] = true;

        }

    }

}

