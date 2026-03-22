_sleepycli_completion() {
    local cur prev=""
    cur="${COMP_WORDS[COMP_CWORD]}"
    if (( COMP_CWORD > 0 )); then
        prev="${COMP_WORDS[COMP_CWORD-1]}"
    fi

    local mode_flags="--now --wake --sleep --from --to"
    local option_flags="--buffer --cycles-min --cycles-max"
    local misc_flags="--good-night --version --help"

    local has_now=0
    local has_wake=0
    local has_sleep=0
    local has_from=0
    local has_to=0
    local has_misc=0

    local word
    for word in "${COMP_WORDS[@]:1:COMP_CWORD-1}"; do
        case "$word" in
            --now) has_now=1 ;;
            --wake) has_wake=1 ;;
            --sleep) has_sleep=1 ;;
            --from) has_from=1 ;;
            --to) has_to=1 ;;
            --good-night|--version|--help) has_misc=1 ;;
        esac
    done

    case "$prev" in
        --wake|--sleep|--from|--to|--buffer|--cycles-min|--cycles-max)
            if [[ "$cur" != -* ]]; then
                COMPREPLY=()
                return
            fi
            ;;
    esac

    if (( has_misc == 1 )); then
        COMPREPLY=()
        return
    fi

    local suggestions=""

    if (( has_now == 0 )) && (( has_wake == 0 )) && (( has_sleep == 0 )) && (( has_from == 0 )) && (( has_to == 0 )); then
        suggestions="$mode_flags $misc_flags"
    elif (( has_now == 1 )) || (( has_wake == 1 )) || (( has_sleep == 1 )); then
        suggestions="$option_flags"
    else
        if (( has_from == 1 )) && (( has_to == 0 )); then
            suggestions="--to"
        elif (( has_to == 1 )) && (( has_from == 0 )); then
            suggestions="--from"
        else
            suggestions="$option_flags"
        fi
    fi

    COMPREPLY=( $(compgen -W "$suggestions" -- "$cur") )
}

complete -F _sleepycli_completion sleepycli