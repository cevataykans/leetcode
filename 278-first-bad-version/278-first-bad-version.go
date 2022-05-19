/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    
    return searchFirstBadVer( 1, n)
}

func searchFirstBadVer( start int, end int) int {
    
    if start == end {
        return start
    }
    
    middle := ( start + end ) / 2
    if !isBadVersion( middle - 1) && isBadVersion( middle) {
        return middle
    }
    
    if isBadVersion( middle) {
        return searchFirstBadVer( start, middle - 1)
    }
    return searchFirstBadVer( middle + 1, end)
}