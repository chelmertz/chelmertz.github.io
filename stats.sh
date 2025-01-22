#!/bin/bash

# Step 1: Trim away the front matter
trim_front_matter() {
    input_file=$1
    awk '/^---$/ {in_frontmatter=!in_frontmatter; next} !in_frontmatter' "$input_file"
}

# Step 2: Count words in the trimmed content
count_words() {
    trimmed_content=$1
    echo "$trimmed_content" | wc -w
}

# Step 3: Collect word counts and generate metrics
generate_metrics() {
    word_counts=("$@")
    total=${#word_counts[@]}

    # Sort word counts to compute median
    sorted_counts=($(printf '%s\n' "${word_counts[@]}" | sort -n))

    # Calculate mean
    total_words=$(IFS=+; echo "$((${word_counts[*]}))")
    mean=$((total_words / total))

    # Calculate median
    if ((total % 2 == 0)); then
        mid=$((total / 2))
        median=$(((sorted_counts[mid-1] + sorted_counts[mid]) / 2))
    else
        median=${sorted_counts[$((total / 2))]}
    fi

    # Get minimum and maximum
    min=${sorted_counts[0]}
    max=${sorted_counts[-1]}

    # Display histogram-like representation
    echo "Word Count Distribution:"
    for count in "${sorted_counts[@]}"; do
        echo -n "$count words: "
        printf '%0.s#' $(seq 1 $((count / 10))) # '#' per 10 words
        echo
    done

    echo
    echo "Total Posts: $total"
    echo "Min Words: $min"
    echo "Mean Words: $mean"
    echo "Median Words: $median"
    echo "Max Words: $max"
}

# Main script logic
main() {
    folder="_posts"
    word_counts=()

    # Check if folder exists
    if [[ ! -d "$folder" ]]; then
        echo "Error: Folder '$folder' not found."
        exit 1
    fi

    for file in "$folder"/*.md; do
        echo "Processing file: $file"

        # Trim front matter
        trimmed_content=$(trim_front_matter "$file")

        # Count words
        word_count=$(count_words "$trimmed_content")
        word_counts+=("$word_count")

        echo "Word count for $file: $word_count"
    done

    # Generate metrics and histogram
    generate_metrics "${word_counts[@]}"
}

# Run the script
main

