
type Props = {
    sortByDate: () => void,
    sortByStars: () => void,
}

export const SortButton = (props: Props) => {
    const { sortByDate, sortByStars } = props;
    return (
        <div className="sortButtons">
            <button onClick={sortByDate}>Sort By Date</button>
            <button onClick={sortByStars}>Sort By Stars</button>
        </div>
    )
  }