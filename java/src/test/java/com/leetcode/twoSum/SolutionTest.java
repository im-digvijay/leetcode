package com.leetcode.twoSum;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.params.provider.Arguments.arguments;

class SolutionTest {
    static Stream<Arguments> inputs() {
        return Stream.of(
                arguments(new int[]{2, 7, 11, 15}, 9, new int[]{0, 1}),
                arguments(new int[]{3, 2, 4}, 6, new int[]{1, 2}),
                arguments(new int[]{3, 3}, 6, new int[]{0, 1})
        );
    }

    @ParameterizedTest
    @MethodSource("inputs")
    void testSolution(int[] nums, int target, int[] expected) {
        Solution solution = new Solution();
        assertArrayEquals(expected, solution.twoSum(nums, target));
    }
}